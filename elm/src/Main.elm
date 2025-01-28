module Main exposing (..)

import Browser
import Html exposing (Html, button, div, form, input, text)
import Html.Attributes exposing (placeholder, value, style, type_)
import Html.Events exposing (onClick, onInput, onSubmit)



-- MODÈLE

type alias Model =
    { userInput : String
    , result : String
    }


init : Model
init =
    { userInput = ""
    , result = ""
    }


-- MESSAGES

type Msg
    = UpdateInput String
    | Submit
    | SubmitForm


-- MISE À JOUR

update : Msg -> Model -> Model
update msg model =
    case msg of
        UpdateInput newInput ->
            { model | userInput = newInput }

        Submit ->
            { model | result = model.userInput }
        

        SubmitForm ->
            { model | result = model.userInput }


-- VUE

view : Model -> Html Msg
view model =
    div [ style "display" "flex"
        , style "flex-direction" "column"
        , style "align-items" "center"
        , style "justify-content" "center"
        , style "height" "100vh"
        , style "background-color" "#f0f0f0"
        ]
        [ form [ onSubmit SubmitForm, style "display" "flex", style "align-items" "center" ]
            [ input
                [ type_ "text"
                , placeholder "Entrez votre texte ici..."
                , value model.userInput
                , onInput UpdateInput
                , style "padding" "10px"
                , style "font-size" "16px"
                , style "border" "1px solid #ccc"
                , style "border-radius" "4px"
                , style "outline" "none"
                ]
                []
            , button
                [ type_ "submit"
                , style "padding" "10px 20px"
                , style "font-size" "16px"
                , style "margin-left" "10px"
                , style "border" "none"
                , style "border-radius" "4px"
                , style "background-color" "#4CAF50"
                , style "color" "white"
                , style "cursor" "pointer"
                ]
                [ text "Entrer" ]
            ]
        , div [ style "margin-top" "20px"
              , style "width" "400px"
              , style "height" "200px"
              , style "border" "2px solid #4CAF50"
              , style "border-radius" "8px"
              , style "padding" "20px"
              , style "background-color" "white"
              , style "box-shadow" "0 4px 8px rgba(0, 0, 0, 0.1)"
              , style "display" "flex"
              , style "align-items" "center"
              , style "justify-content" "center"
              ]
            [ text model.result ]
        ]


-- PROGRAMME PRINCIPAL

main : Program () Model Msg
main =
    Browser.sandbox { init = init, update = update, view = view }
