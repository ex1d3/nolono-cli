package commands

import (
	"errors"
	"fmt"
	"jasmine-cli/executor/collections"
	"jasmine-cli/executor/commands/internal_errors"
	"jasmine-cli/storage"
)

// example args silces: [tx 1] or [src adam]
func Del(args []string) ([]string, error) {
	if len(args) != 2 {
		return []string{}, errors.New(
			internal_errors.InvalidArgsCount("del", "2", len(args)),
		)
	}

	collection := args[0]
	target := args[1]

	switch collection {
	case collections.SRC:
		{
			return delSrc(target)
		}
	case collections.TX:
		{
			return delTx(target)
		}
	default:
		{
			return []string{}, errors.New(
				internal_errors.InvalidCollection(collection),
			)
		}
	}
}

func delSrc(target string) ([]string, error) {
	srcStorage := storage.Src

	if target == "*" {
		srcStorage.Unload()

		return []string{}, nil
	}

	if srcStorage.Get(target) == nil {
		return []string{}, errors.New(
			invalidTargetForCollection(target, collections.SRC),
		)
	}

	srcStorage.Delete(target)

	return []string{storage.NullStoragePointer(target)}, nil
}

func delTx(target string) ([]string, error) {
	txStorage := storage.Tx

	if target == "*" {
		txStorage.Unload()

		return []string{}, nil
	}

	if tx := txStorage.Get(target); tx == nil {
		return []string{}, errors.New(
			invalidTargetForCollection(target, collections.TX),
		)
	}

	txStorage.Delete(target)

	return []string{storage.NullStoragePointer(target)}, nil
}

func invalidTargetForCollection(target string, collection string) string {
	return fmt.Sprintf(
		"invalid target for collection '%s' (%s)",
		collection,
		target,
	)
}
