# Project for 02-601
This is the final project for 02-601 programming for scientists. In this project, firstly, we build models for simulating the behaviour of hysarum polycephalum molds to find the shortest path between 2 food spots in the maze. Then this model is generalized into finding the shortest path between 2 cities in the road map of China. Later, the multi-agent model is ultilized for simulating the behaviour of hysarum polycephalum molds to seek the shortest paths among several food spots in the 2D board. Other parameters such as wind, light, good food and bad food are added to simulate the behaviour in a more practical way. Finally, a user interface is built to simulate the process interactively and simultanously.

## Getting Started

The instructions below will help you to download the code and run on your local machine to generate results. You'll need the go package and compiler such as atom for running go code.

### Installing
```
git clone https://github.com/wwwsl0727/Programming_Project.
```

## Part 1: The shortest path between two spots in a 2D board.

## Running the model
```
go build SPinMaze
```

For simulating the simple problem in Physarum polycephalum maze experiment, mode="maze". For simulating the shortest path between two cities in China from the road map, mode ="transport".

We only add light to the maze solving problem of Physarum polycephalum molds. If there is light, isLight=true, delete one tube in the original maze. Else, islight=false.

The fileName is for the name of the output GIF.

```
./SPinMaze mode isLight fileName
```
## Part 2: Multi-Agent

## Part 3: interface
