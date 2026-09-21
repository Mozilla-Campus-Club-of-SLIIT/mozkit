#!/usr/bin/env bash
set -e

# Ensure gum retains ANSI colors when captured into variables
export CLICOLOR_FORCE=1

# Mozkit Theme Highlight Color
COLOR_ORANGE="${COLOR_ORANGE:-#ff4b0c}"

# Common Moxy prompt label
MOXY=$(gum style --foreground "$COLOR_ORANGE" --bold "Moxy: ")

# Spin helper: runs gum spin and drains unconsumed terminal capability responses silently using read -rs without changing global stty state
spin() {
  local title="$1"
  shift
  gum spin --spinner dot --spinner.foreground "$COLOR_ORANGE" --title "$title" -- "$@" < /dev/null
}

# Greeting
gum join "$MOXY" "$(gum style 'Hey! Ready to get some work done today?')"
gum style "Let's set up a brand new branch so you can just focus on the fun part."
echo ""

# Step 1: Check GitHub CLI Authentication with spinner
TMP_AUTH=$(mktemp)
spin "Checking your GitHub CLI login..." \
  bash -c 'gh auth status -a --json hosts --jq ".hosts[\"github.com\"][].login" > "'"$TMP_AUTH"'" 2>/dev/null || true'

LOGIN_USER=$(cat "$TMP_AUTH" 2>/dev/null || true)
rm -f "$TMP_AUTH"

if [ -z "$LOGIN_USER" ]; then
  AUTH_CMD=$(gum style --foreground "$COLOR_ORANGE" --bold "gh auth login")
  gum join "$MOXY" "$(gum style 'Wait a second... I checked, but you'\''re not logged into GitHub CLI yet!')"
  gum join "$(gum style 'Could you run ')" "$AUTH_CMD" "$(gum style ' real quick? I'\''ll wait right here.')"
  exit 1
fi

export GH_USER="$LOGIN_USER"
USER_HL=$(gum style --foreground "$COLOR_ORANGE" --bold "@$GH_USER")
gum join "$MOXY" "$(gum style 'Nice to see you, ')" "$USER_HL" "$(gum style '! Let'\''s get things rolling.')"

# Step 2: Ensure on main branch
CURRENT_BRANCH=$(git branch --show-current)

if [ "$CURRENT_BRANCH" != "main" ]; then
  CUR_HL=$(gum style --foreground "$COLOR_ORANGE" --bold "$CURRENT_BRANCH")
  MAIN_HL=$(gum style --foreground "$COLOR_ORANGE" --bold "main")
  gum join "$MOXY" "$(gum style 'You'\''re currently on ')" "$CUR_HL" "$(gum style '. Let'\''s hop over to ')" "$MAIN_HL" "$(gum style ' first so we start clean...')"
  
  if ! git checkout main >/dev/null 2>&1; then
    echo ""
    gum join "$MOXY" "$(gum style 'Uh oh... Git got a little stubborn switching to ')" "$MAIN_HL" "$(gum style '.')"
    gum style "Looks like there might be some uncommitted changes or files in the way."
    gum style "Mind saving or stashing them real quick so we can try again?"
    exit 1
  fi
fi

# Step 3: Pull latest changes with a spinner
if ! spin "Syncing up with origin/main..." \
  bash -c 'git pull >/dev/null 2>&1'; then
  echo ""
  gum join "$MOXY" "$(gum style 'Oh no, git pull ran into a little snag!')"
  gum style "Seems like your local branch and origin aren't syncing up right now."
  gum style "Could you check your network or branch status so we can get back on track?"
  exit 1
fi

# Step 4: Fetch assigned issues with a spinner and let user choose
TMP_ISSUES=$(mktemp)
spin "Checking what issues you're working on..." \
  bash -c "gh issue list --state open --assignee '$GH_USER' --json number,title --jq '.[] | \"#\(.number) - \(.title)\"' > '$TMP_ISSUES' 2>/dev/null || true"

ISSUES_RAW=$(cat "$TMP_ISSUES" 2>/dev/null || true)
rm -f "$TMP_ISSUES"

if [ -z "$ISSUES_RAW" ]; then
  echo ""
  USER_HL=$(gum style --foreground "$COLOR_ORANGE" --bold "@$GH_USER")
  gum join "$MOXY" "$(gum style 'Hmm, I looked around GitHub but couldn'\''t find any open issues assigned to ')" "$USER_HL" "$(gum style '!')"
  gum style "Go ahead and grab an issue on GitHub or ask a maintainer to assign one to you, then let's do this!"
  exit 0
fi

# Line break before prompt
echo ""
gum join "$MOXY" "$(gum style 'Which issue are we tackling today?')"

# Explicitly pass --header="" so gum doesn't output the default "Choose:" prompt
SELECTED_ISSUE=$(echo "$ISSUES_RAW" | gum choose --limit 1 \
  --header="" \
  --cursor.foreground "$COLOR_ORANGE" \
  --selected.foreground "$COLOR_ORANGE")

if [ -z "$SELECTED_ISSUE" ]; then
  echo ""
  gum join "$MOXY" "$(gum style 'No problem! We can always pick this up later whenever you'\''re ready. 😊')"
  exit 0
fi

# Extract issue number from selection (e.g. '#3 - ...' -> '3')
ISSUE_NUM=$(echo "$SELECTED_ISSUE" | sed -E 's/^#?([0-9]+).*/\1/')

# Step 5: Create branch with spinner
TMP_ERR=$(mktemp)
if spin "Cooking up your new branch for issue #$ISSUE_NUM..." \
  bash -c "gh issue develop -c '$ISSUE_NUM' >/dev/null 2>'$TMP_ERR'"; then
  rm -f "$TMP_ERR"
  echo ""
  ISSUE_HL=$(gum style --foreground "$COLOR_ORANGE" --bold "#$ISSUE_NUM")
  gum join "$MOXY" "$(gum style 'Yay, all done! 🎉 Your branch for issue ')" "$ISSUE_HL" "$(gum style ' is checked out and ready for you!')"
  gum style "Go do your thing, I know you're gonna crush it! ✨💻"
else
  echo ""
  gum join "$MOXY" "$(gum style 'Oh no, something went sideways creating the branch...')"
  if [ -s "$TMP_ERR" ]; then
    cat "$TMP_ERR"
  fi
  rm -f "$TMP_ERR"
  gum style "Take a peek at the error output above to see what happened, okay?"
  exit 1
fi
