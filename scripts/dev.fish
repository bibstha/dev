# This fish function need to exist to do `dev cd something`
#
# A binary launched for a shell cannot change the directory of the shell.  Such binary is a separate process and can
# change the binary's own cwd, but as soon as it exits, the parent process (SHELL)'s cwd remains the same as it was.
# Shell's binary can only be changed by the shell's function.
# * https://stackoverflow.com/q/255414/84143
# * https://stackoverflow.com/q/53984853/84143
#
# This function acts as a shim (proxy) to the actual dev binary
function dev -a cmd
  set cmdArgv $argv[2..-1]
  set devBinary (which dev)

  switch $cmd
    case 'cd'
      set newDir ($devBinary cd $cmdArgv)
      if test $status = 0
        cd $newDir
      else
        return $status
      end
    case '*'
      $devBinary $cmd $cmdArgv
  end
end
