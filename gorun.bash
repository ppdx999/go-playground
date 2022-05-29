#!/bin/bash

# completion for bash

_func() {
	local cur prev words cword split
	_init_completion -s || return

	case $cur in
	*)
		_filedir
		return
		;;
	esac
}

complete -F _func "./${1#.*}"
