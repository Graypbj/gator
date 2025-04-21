package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Graypbj/gator/internal/config"
	"github.com/Graypbj/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	userID := uuid.New()

	now := time.Now()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        userID,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			fmt.Printf("User %s already exists\n", name)
			os.Exit(1)
		}
		return fmt.Errorf("error creating user: %w", err)
	}

	s.cfg.CurrentUserName = name
	if err := config.Write(*s.cfg); err != nil {
		return fmt.Errorf("error writing config: %w", err)
	}

	fmt.Printf("User %s created successfully!\n", name)
	printUser(user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	// Check if the user exists in the database
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			// User doesn't exist, exit with code 1
			fmt.Printf("User %s does not exist\n", name)
			os.Exit(1)
		}
		return fmt.Errorf("error looking up user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed name: %v\n", feed.FeedName)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:     %v\n", user.ID)
	fmt.Printf(" * Name:   %v\n", user.Name)
}
