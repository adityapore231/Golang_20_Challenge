
#writing a programm to generate heart using turtle
# write happy valentine's day in the middle of the heart
# write your name in the bottom of the heart

import turtle
import time

#creating a turtle object
t = turtle.Turtle()
t.speed(10)
t.hideturtle()

#creating a screen object
s = turtle.Screen()
s.bgcolor("black")
s.title("Happy Valentine's Day")

#function to draw a heart
def draw_heart():
    t.begin_fill()
    t.fillcolor("red")
    t.left(140)
    t.forward(180)
    t.circle(-90, 200)
    
    #draw left curve
    t.left(120)
    t.circle(-90, 200)
    t.forward(180)
    t.end_fill()

#function to write happy valentine's day
def write_text():
    time.sleep(2)
    t.penup()
    t.goto(-30, 0)
    t.pendown()
    t.color("pink")
    t.write("Happy Valentine's Day", font=("Verdana", 12, "bold"))
    t.hideturtle()

#function to write your name
def write_name():
    time.sleep(2)
    t.penup()
    t.goto(-30, -100)
    t.pendown()
    t.color("pink")
    t.write("Shruti", font=("Verdana", 12, "bold"))
    t.hideturtle()

#drawing heart
draw_heart()
write_text()
write_name()

#to hold the window
turtle.mainloop()