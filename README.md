# The Dining Philosophers in Go

This is our (silk@itu.dk, sefu@itu.dk, milo@itu.dk, krwe@itu.dk) solutions to the first mini-project of the course Distributed Systems, BSc at the IT University of Copenhagen, spring 2021.

## The assignment is as follows:

The Dining Philosophers is a well known problem in Computer Science that concerns concurrency. At a dining round table there are five philosophers who are supposed to have dinner. Philosophers are kind of special and while they have dinner, they either *eat* their food or *think* about something. Although they can think at any time, in order to be able to eat they must get hold of two forks (the food is very special and cannot be handled with one fork). Unfortunately, there are only five forks at the table, each of them uniquely placed between two philosophers (the table is round, there is exactly one fork between any two philosophers -- each philosopher can only reach the two forks that are nearby). As a consequence, it is never the case that all philosophers can eat at the same time (max two at the time). This problem is interesting because, depending on how they decide to pick the forks, the philosopher may reach a deadlock. The goal of this project is to implement the dining philosophers problem in Go, with the following requirements:

- Each philosopher and each fork must have its own goroutine (you must use concurrency) -- please use separate files each.

- The system must be designed in a way that does no lead to a deadlock

- Each philosopher must include two channels (one for input and one for output, both usable from the outside) through which it is possible to make queries on the state of the philosopher (number of times eaten, eating or thinking)

- Each fork must include two channels (one for input and one for output, both usable from outside) through which it is possible to make queries on the state of the fork (number of times used, in use or free)
