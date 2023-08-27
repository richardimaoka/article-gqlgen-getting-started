#!/bin/sh

PWD=$(pwd)

SESSION=$(basename "$PWD")

tmux new-session -s "$SESSION" -d
tmux send-keys -t "$SESSION:0.0" 'gow -e=graphqls run github.com/99designs/gqlgen generate' C-m

tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane'
tmux send-keys -t "$SESSION:0.1" 'gow -e=go,json run .' C-m

tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane'
tmux send-keys -t "$SESSION:0.2" 'code gqlgen' C-m

tmux attach -t "$SESSION"
