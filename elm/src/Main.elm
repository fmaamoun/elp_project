module Main exposing (main)

import Browser
import Html exposing (Html, button, div, form, input, text)
import Html.Attributes exposing (class, placeholder, type_, value)
import Html.Events exposing (onInput, onSubmit)
import TurtleParser exposing (read, Command(..), commandsToString)
import TurtleDrawing exposing (display)


-- MODEL
type alias Model =
    { userInput : String
    , result : String
    , commands : Maybe (List Command)
    }


init : Model
init =
    { userInput = ""
    , result = ""
    , commands = Nothing
    }


-- UPDATE
type Msg
    = UpdateInput String
    | SubmitForm


update : Msg -> Model -> Model
update msg model =
    case msg of
        UpdateInput input ->
            { model | userInput = input }

        SubmitForm ->
            case read model.userInput of
                Ok cmds ->
                    { model
                        | result = "Parsed: " ++ commandsToString cmds
                        , commands = Just cmds
                    }

                Err err ->
                    { model
                        | result = "Parsing error: " ++ err
                        , commands = Nothing
                    }


-- VIEW
view : Model -> Html Msg
view model =
    div [ class "container" ]
        [ form [ class "form-container", onSubmit SubmitForm ]
            [ input
                [ class "input-field"
                , type_ "text"
                , placeholder "Enter commands..."
                , value model.userInput
                , onInput UpdateInput
                ]
                []
            , button [ class "submit-button", type_ "submit" ] [ text "Submit" ]
            ]
        , div [ class "result-box" ]
            (case model.commands of
                Just cmds ->
                    [ display cmds ]
                Nothing ->
                    [ text model.result ]
            )
        ]


main : Program () Model Msg
main =
    Browser.sandbox { init = init, update = update, view = view }
