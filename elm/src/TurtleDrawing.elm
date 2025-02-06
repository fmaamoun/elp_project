module TurtleDrawing exposing (displayPartial)

import Svg exposing (Svg, svg, polyline)
import Svg.Attributes exposing (points, viewBox, width, height, stroke, strokeWidth, fill)
import List
import TurtleParser exposing (Command(..))
import Basics exposing (cos, sin, pi)
import String


-- A point in 2D space.
type alias Point =
    { x : Float, y : Float }


-- Turtle state: current position, angle (in degrees) and path.
type alias State =
    { pos : Point, angle : Float, path : List Point }


-- Convert degrees to radians.
toRadians : Float -> Float
toRadians deg =
    deg * pi / 180


-- Move the turtle forward by a given distance.
move : Float -> State -> State
move dist state =
    let
        rad = toRadians state.angle
        newPos =
            { x = state.pos.x + dist * cos rad
            , y = state.pos.y + dist * sin rad
            }
    in
    { state | pos = newPos, path = state.path ++ [ newPos ] }


-- Process a single turtle command.
processCommand : State -> TurtleParser.Command -> State
processCommand state cmd =
    case cmd of
        TurtleParser.Forward n ->
            move (toFloat n) state

        TurtleParser.Back n ->
            move (toFloat (-n)) state

        TurtleParser.Left n ->
            { state | angle = state.angle + toFloat n }

        TurtleParser.Right n ->
            { state | angle = state.angle - toFloat n }

        TurtleParser.Repeat n cmds ->
            iterate (processCommands cmds) n state


-- Helper: iterate a function n times.
iterate : (State -> State) -> Int -> State -> State
iterate f n state =
    if n <= 0 then
        state
    else
        iterate f (n - 1) (f state)


-- Process a list of commands sequentially.
processCommands : List TurtleParser.Command -> State -> State
processCommands cmds state =
    List.foldl (\cmd st -> processCommand st cmd) state cmds


-- Compute the turtle's path up to a given step.
computePath : List TurtleParser.Command -> Int -> List Point
computePath cmds stepCount =
    let
        initState =
            { pos = { x = 0, y = 0 }
            , angle = 0
            , path = [ { x = 0, y = 0 } ]
            }
        limitedCmds = List.take stepCount cmds
    in
    (processCommands limitedCmds initState).path


-- Convert a list of points into a string for the SVG "points" attribute.
pointsStr : List Point -> String
pointsStr pts =
    pts
        |> List.map (\p -> String.fromFloat p.x ++ "," ++ String.fromFloat p.y)
        |> String.join " "


-- Generate the SVG drawing progressively.
displayPartial : List TurtleParser.Command -> Int -> Svg msg
displayPartial cmds stepCount =
    let
        pts = computePath cmds stepCount
    in
    svg [ width "400", height "400", viewBox "-200 -200 400 400" ]
        [ polyline [ points (pointsStr pts), stroke "black", strokeWidth "2", fill "none" ] [] ]
