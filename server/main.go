package main

/*
	Steps
	Server takes conns the same way
	a safe map is created, methods to add and remove are added
	Map adds when a connection is created
	Map removes when a connection is closed and removes all when a transaction is made
	A connection is checked for being closed before transaction, by setting deadline to 10 miliseconds, reading 1 byte and checking if error is EOF
	The array of conns created by the map is trimmed down of his closed connections
	A buffer of size 1MB ( 1 << 20 ?)is created to hold data
	Read from sender to buff
	Create a io.MultiWriter to write to all conns(use spread operator)
	After that all connections are closed and the Map value is set to nil
	The Client must have a flag to automatically restart the connection after a file is transfered
	Both the client and the server must gracefully close all connections when a key is pressed
	A report object must be created and pushed to DB , regardless of state
*/
