module TcTurtleParser exposing (Command(..), read, commandsToString)

import Parser exposing
    ( Parser
    , run
    , oneOf
    , symbol
    , spaces
    , succeed
    , (|.)
    , (|=)
    , float
    , andThen
    , lazy
    )
import String exposing (fromInt)


{-|
    1) TYPE DE COMMANDE
-}
type Command
    = Forward Int
    | Back Int
    | Left Int
    | Right Int
    | Repeat Int (List Command)


{-|
    2) AFFICHAGE
-}
commandsToString : List Command -> String
commandsToString cmds =
    cmds
        |> List.map commandToString
        |> String.join ", "


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
            "Repeat " ++ fromInt n
                ++ " ["
                ++ commandsToString subcmds
                ++ "]"


{-|
    3) PARSER D’UN ENTIER
    On lit un float, qu’on convertit en Int
-}
intParser : Parser Int
intParser =
    float
        |> andThen (\f ->
            succeed (round f)
        )


{-|
    4) CHACUNE DES COMMANDES SIMPLES
-}
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


{-|
    5) PARSER "REPEAT"
    Il se réfère à `commandsParser ()`, donc la liste de commandes
-}
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


{-|
    6) "commandParser" DEVIENT UNE FONCTION
    pour éviter la boucle de dépendances
-}
commandParser : Parser Command
commandParser =
    oneOf
        [ forwardParser
        , backParser
        , leftParser
        , rightParser
        , repeatParser
        ]


{-|
    7) LISTE DE COMMANDES
    Elle se réfère à `commandParser ()`
-}
commandsParser : Parser (List Command)
commandsParser =
    succeed (::)
        |= lazy (\_ -> commandParser)
        |= lazy (\_ -> manyCommandsParser)


manyCommandsParser : Parser (List Command)
manyCommandsParser =
    oneOf
        [ succeed (\c tail -> c :: tail)
            |. symbol ","
            |. spaces
            |= lazy (\_ -> commandParser)
            |= lazy (\_ -> manyCommandsParser)
        , succeed []
        ]


programParser : Parser (List Command)
programParser =
    succeed identity
        |. symbol "["
        |. spaces
        |= commandsParser
        |. spaces
        |. symbol "]"


read : String -> Result String (List Command)
read input =
    case run programParser input of
        Ok cmds ->
            Ok cmds

        Err _ ->
            Err "error"