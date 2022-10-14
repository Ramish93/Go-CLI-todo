# Golang CLI Todo List:

- This golang project is developed mainly with native library.
- To run the project download and unzip files, open in IDE run **go mod tidy**.
- Now run the command **go build ./cmd/todo** .
- To add a todo run **./todo -add < your todo >**
- After adding a few todos see the list in CLI with command **./todo -list**
- To delete a todo run **./todo -delete=2 todo number**
- To complete a todo run **./todo -complete=1 < your todo >**
- The data is stored in a json file in project. you can manually edit the data but a single mistake will currept the data.

![image of preview](./preview.png?raw=true "Hero preview image")
