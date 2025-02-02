module TurtleParser exposing (Command(..), read, commandsToString)

import Parser exposing
    ( Parser, run, oneOf, symbol, spaces, succeed, (|.), (|=), float, andThen, lazy )
import String exposing (fromInt)


-- COMMAND TYPE DEFINITION
type Command
    = Forward Int
    | Back Int
    | Left Int
    | Right Int
    | Repeat Int (List Command)


-- Convert a list of commands to a string.
commandsToString : List Command -> String
commandsToString cmds =
    String.join ", " (List.map commandToString cmds)


commandToString : Command -> String
commandToString cmd =
    case cmd of
        Forward n ->
            "Forward " ++ fromInt n

        Back n ->
            "Back " ++ fromInt n

        Left n ->
            "Left " ++ fromInt n

        Right n ->
            "Right " ++ fromInt n

        Repeat n subcmds ->
            "Repeat " ++ fromInt n ++ " [" ++ commandsToString subcmds ++ "]"


-- PARSER FOR AN INTEGER (reads a float then rounds it)
intParser : Parser Int
intParser =
    float |> andThen (\f -> succeed (round f))


-- SIMPLE COMMAND PARSERS
forwardParser : Parser Command
forwardParser =
    succeed Forward
        |. symbol "Forward"
        |. spaces
        |= intParser


backParser : Parser Command
backParser =
    succeed Back
        |. symbol "Back"
        |. spaces
        |= intParser


leftParser : Parser Command
leftParser =
    succeed Left
        |. symbol "Left"
        |. spaces
        |= intParser


rightParser : Parser Command
rightParser =
    succeed Right
        |. symbol "Right"
        |. spaces
        |= intParser


repeatParser : Parser Command
repeatParser =
    succeed Repeat
        |. symbol "Repeat"
        |. spaces
        |= intParser
        |. spaces
        |. symbol "["
        |. spaces
        |= lazy (\_ -> commandsParser)
        |. spaces
        |. symbol "]"


commandParser : Parser Command
commandParser =
    oneOf [ forwardParser, backParser, leftParser, rightParser, repeatParser ]


-- PARSER FOR A LIST OF COMMANDS SEPARATED BY COMMAS
commandsParser : Parser (List Command)
commandsParser =
    succeed (::)
        |= lazy (\_ -> commandParser)
        |= manyCommandsParser


manyCommandsParser : Parser (List Command)
manyCommandsParser =
    oneOf
        [ succeed (\cmd tail -> cmd :: tail)
            |. symbol ","
            |. spaces
            |= commandParser
            |= lazy (\_ -> manyCommandsParser)
        , succeed []
        ]



-- A PROGRAM IS A LIST OF COMMANDS ENCLOSED IN BRACKETS
programParser : Parser (List Command)
programParser =
    succeed identity
        |. symbol "["
        |. spaces
        |= commandsParser
        |. spaces
        |. symbol "]"


-- READ FUNCTION: Parses input into a list of commands.
read : String -> Result String (List Command)
read input =
    case run programParser input of
        Ok cmds ->
            Ok cmds

        Err _ ->
            Err "Error parsing input"
