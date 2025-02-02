# TcTurtle Web Application

This project is an Elm web application that visualizes turtle graphics based on user-entered commands. The commands follow TcTurtle syntax and support `Forward`, `Back`, `Left`, `Right`, and `Repeat` instructions.


## 🚀 How to Run the Project

Navigate to the `build/` folder and open `index.html` in a browser.


## 🏗️ How to Build the Project

### 1️⃣ Prerequisites

- Install [Elm](https://elm-lang.org/)

### 2️⃣ Build the Elm Application

Run the following command from the **elm/** directory:

```sh
elm make src/Main.elm --optimize --output=build/index.js
```

This compiles the Elm code into optimized JavaScript.

### 3️⃣ Open the Application

Navigate to the `build/` folder and open `index.html` in a browser.


## 📜 How It Works

1️⃣ **User Input**: Enter commands into the text field.  
2️⃣ **Parsing**: The input is parsed using `TurtleParser.elm`.  
3️⃣ **Drawing**: If parsing is successful, `TurtleDrawing.elm` generates an SVG.  
4️⃣ **Display**: The parsed result (or error message) and the SVG drawing appear inside the **green-bordered result box**.


## 🖊️ Example Commands

Try these commands in the input field:

✅ **Circle pattern**:
```
[Repeat 360 [Right 1, Forward 1]]
```

✅ **Square with a line**:
```
[Forward 100, Repeat 4 [Forward 50, Left 90], Forward 100]
```

✅ **Star pattern**:
```
[Repeat 36 [Right 10, Repeat 8 [Forward 25, Left 45]]]
```

✅ **Complex nested loops**:
```
[Repeat 8 [Left 45, Repeat 6 [Repeat 90 [Forward 1, Left 2], Left 90]]]
```