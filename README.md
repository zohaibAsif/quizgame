# Quizgame
Gophercises Project 1: Quiz Game

# Commands

-to run with default values (unshuffled quiz from problems.csv file with a time limit of 30 seconds).

  _go run quiz_app.go_
  
  
-to run with custom values (shuffled quiz from abc.csv file with a time limit of 10 seconds).
  
  _go run quiz_app -csv=abc.csv -limit=10 -shuffle=true_
  
# Flags
  
_-csv=filename.csv_

_-limit=time limit for quiz_

_-shuffle=true or false_
