package warehouse

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search warehouse for pallets",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run:  search,
}

func init() {
	warehouseCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func search(cmd *cobra.Command, args []string) {
	arg := args[0]
	parts := strings.Split(arg, ":")
	if len(parts) != 2 {
		_ = fmt.Errorf("invalid argument format. Must be <prefix>:<slug>")
		os.Exit(1)
	}
	prefix := parts[0]
	slug := parts[1]
	url := getGitRepoURL(prefix, slug)

	if url == "" {
		_ = fmt.Errorf("unsupported prefix '%s'", prefix)
		os.Exit(1)
	}

	fmt.Println(url)

	// Check if the repository exists remotely
	_, err := git.PlainClone(slug, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil && err != transport.ErrRepositoryNotFound {
		_ = fmt.Errorf(err.Error())
		os.Exit(1)
	} else if err == transport.ErrRepositoryNotFound {
		_ = fmt.Errorf("repository not found at %s", url)
		os.Exit(1)
	}

	fmt.Printf("Repository cloned successfully from %s\n", url)
	return
}

func getGitRepoURL(prefix string, slug string) string {
	var website string
	switch prefix {
	case "gh":
		website = "https://github.com/"
	case "gl":
		website = "https://gitlab.com/"
	default:
		return ""
	}
	return website + slug
}
