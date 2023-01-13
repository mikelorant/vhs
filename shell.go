package main

// Supported shells of VHS
const (
	bash       = "bash"
	cmdexe     = "cmd"
	fish       = "fish"
	powershell = "powershell"
	pwsh       = "pwsh"
	zsh        = "zsh"
)

// Shell is a type that contains a prompt and the command to set up the shell.
type Shell struct {
	Prompt  string
	Command string
}

// Shells contains a mapping from shell names to their Shell struct.
var Shells = map[string]Shell{
	bash: {
		Prompt:  "\\[\\e[38;2;90;86;224m\\]> \\[\\e[0m\\]",
		Command: ` clear; PS1="%s" bash --login --norc --noprofile; set +o history; clear;`,
	},
	zsh: {
		Prompt:  `%F{#5B56E0}> %F{reset_color}`,
		Command: ` clear; PROMPT="%s" zsh --login --histnostore --no-rcs; clear`,
	},
	fish: {
		Prompt:  `function fish_prompt; echo -e "$(set_color 5B56E0)> $(set_color normal)"; end`,
		Command: ` clear; fish --login --no-config --private -C 'function fish_greeting; end' -C '%s'`,
	},
	powershell: {
		Prompt:  "Function prompt {Write-Host \\\"> \\\" -ForegroundColor Blue -NoNewLine; return \\\"`0\\\" }",
		Command: ` clear; powershell -NoLogo -NoExit -Command 'Set-PSReadLineOption -HistorySaveStyle SaveNothing; %s'`,
	},
	pwsh: {
		Prompt:  "Function prompt {Write-Host \\\"> \\\" -ForegroundColor Blue -NoNewLine; return \\\"`0\\\" }",
		Command: ` clear; pwsh -Login -NoLogo -NoExit -Command 'Set-PSReadLineOption -HistorySaveStyle SaveNothing; %s'`,
	},
	cmdexe: {
		Prompt:  "$g",
		Command: ` cls && set prompt=%s && cls`,
	},
}
