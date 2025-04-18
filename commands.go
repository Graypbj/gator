package main

import "errors"

// Command holds the name of the command and any number of args passed in
type command struct {
	Name string
	Args []string
}

// Commands holds a map of all posible commands and the signature of the command must include: *state, command
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// Register function simply takes in commands and names to then be stored in the commands struct
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// Run function calls the command
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
