module Main exposing (main)

import Browser
import Html exposing (Html, button, div, form, input, text)
import Html.Attributes exposing (class, placeholder, type_, value, disabled)
import Html.Events exposing (onClick, onInput, onSubmit)
import TurtleParser exposing (read, Command(..), commandsToString)
import TurtleDrawing exposing (displayPartial)
import Process
import Task exposing (perform)


-- MODEL
type alias Model =
    { userInput : String
    , result : String
    , commands : Maybe (List Command)
    , currentStep : Int
    , isPlaying : Bool
    }


init : Model
init =
    { userInput = ""
    , result = ""
    , commands = Nothing
    , currentStep = 0
    , isPlaying = False
    }


-- UPDATE
type Msg
    = UpdateInput String
    | SubmitForm
    | Tick
    | StartDrawing
    | ResetDrawing


flattenCommands : List Command -> List Command
flattenCommands cmds =
    List.concatMap flattenCommand cmds


flattenCommand : Command -> List Command
flattenCommand cmd =
    case cmd of
        Repeat n cmds ->
            List.concat (List.repeat n (flattenCommands cmds))

        _ ->
            [ cmd ]



update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        UpdateInput input ->
            ( { model | userInput = input }, Cmd.none )

        SubmitForm ->
            case read model.userInput of
                Ok cmds ->
                    let
                        flattenedCmds = flattenCommands cmds
                    in
                    ( { model
                        | result = "Parsed: " ++ commandsToString cmds
                        , commands = Just flattenedCmds
                        , currentStep = 0
                        , isPlaying = False
                      }
                    , Cmd.none
                    )

                Err err ->
                    ( { model | result = "Parsing error: " ++ err, commands = Nothing }, Cmd.none )

        StartDrawing ->
            ( { model | isPlaying = True }
            , Process.sleep 0.000005 |> perform (\_ -> Tick)
            )

        Tick ->
            case model.commands of
                Just cmds ->
                    if model.currentStep < List.length cmds then
                        ( { model | currentStep = model.currentStep + 1 }
                        , Process.sleep 0.000005 |> perform (\_ -> Tick)
                        )
                    else
                        ( { model | isPlaying = False }, Cmd.none )

                Nothing ->
                    ( model, Cmd.none )

        ResetDrawing ->
            ( { model | currentStep = 0, isPlaying = False }, Cmd.none )


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
        , div [ class "controls" ]
            [ button [ class "play-button", onClick StartDrawing, disabled (model.isPlaying || model.commands == Nothing) ] [ text "Play" ]
            , button [ class "reset-button", onClick ResetDrawing, disabled (model.currentStep == 0) ] [ text "Reset" ]
            ]
        , div [ class "result-box" ]
            (case model.commands of
                Just cmds ->
                    [ displayPartial cmds model.currentStep ]
                Nothing ->
                    [ text model.result ]
            )
        ]


main : Program () Model Msg
main =
    Browser.element { init = \_ -> (init, Cmd.none), update = update, view = view, subscriptions = \_ -> Sub.none }
