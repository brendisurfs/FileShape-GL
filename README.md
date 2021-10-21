# FILESHAPE
by Brendan Prednis

**NOTE: his is a work in progress, lots of changes will be made.**

## what the heck is it

This is what happens when I really want to learn something but can't concentrate on it by any practical sense. 
The only solution is creating a nonsense project, and that is exactly what we have here. 

- What I want to learn? Go + OpenGL 
- What gave me a headache? OpenGL 

- What looks kind of like float32 slices?   byte slices. 

- What are geometry points in OpenGL made of? float32 slices. 

So I decided to combine a few project ideas into one, and this fantastic learning project was born. 

It takes a file upload, reads it to []byte, converts that []byte to []float32, and then we pipe that into OpenGL geometry, hopefully giving us some really messed up geometry on the screen.

This kind of combines alot of aspects I want to brush up on with Go:
- file handling
- type conversion 
- sorting logic
