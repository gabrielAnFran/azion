package edge_functions_instances

import "errors"

var (
	ErrorGetFunctions           = errors.New("failed to get the Edge Functions Instances: %s. Check your settings and try again. If the error persists, contact Azion support")
	ErrorMissingArgumentsDelete = errors.New("Required flags are missing. You must supply application-id and instance-id as arguments. Run 'azioncli <command> <subcommand> --help' command to display more information and try again")
	ErrorFailToDeleteFuncInst   = errors.New("Failed to delete the Edge Function Instance: %s. Check your settings and try again. If the error persists, contact Azion support")
	ErrorMandatoryFlags         = errors.New("One or more required flags are missing. You must provide the --application-id and --instance-id flags. Run the command 'azioncli describe <subcommand> --help' to display more information and try again.")
	ErrorGetEdgeFuncInstances   = errors.New("Failed to describe the Edge Function Instance: %s. Check your settings and try again. If the error persists, contact Azion support.")
	ErrorUpdateFuncInstance     = errors.New("Failed to update the Edge Function Instance: %s. Check your settings and try again. If the error persists, contact Azion support.")
	ErrorMandatoryUpdateFlags   = errors.New("Required flags are missing. You must provide application-id, instance-id and function-id flags when the --in flag is not provided. Run the command 'azioncli <command> <subcommand> --help' to display more information and try again.")
	ErrorMandatoryUpdateFlagsIn = errors.New("Required flags are missing. You must provide application-id and instance-id flags when the --in flag is provided. Run the command 'azioncli <command> <subcommand> --help' to display more information and try again.")
)
