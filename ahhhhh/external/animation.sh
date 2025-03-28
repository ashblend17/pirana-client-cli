#!/usr/bin/env bash

# shellcheck disable=SC2034 # https://github.com/koalaman/shellcheck/wiki/SC2034

# Load in the functions and animations
source ./ahhhhh/external/bash_loading_animations.sh
# Run BLA::stop_loading_animation if the script is interrupted
trap BLA::stop_loading_animation SIGINT SIGTERM EXIT

demo_loading_animation() {
  BLA_active_loading_animation=( "${@}" )
  # Extract the delay between each frame from the active_loading_animation array
  BLA_loading_animation_frame_interval="${BLA_active_loading_animation[0]}"
  # Sleep long enough that all frames are showed
  # substract 1 to the number of frames to account for index [0]
  demo_duration=$( echo "${BLA_active_loading_animation[0]} * ( ${#BLA_active_loading_animation[@]} - 1 )" | bc )
  # Make sure each animation is shown for at least 3 seconds
  if [[ $( echo "if (${demo_duration} < 3) 0 else 1" | bc ) -eq 0 ]] ; then
    demo_duration=3
  fi
  unset "BLA_active_loading_animation[0]"
  echo
  BLA::play_loading_animation_loop &
  BLA_loading_animation_pid="${!}"
  sleep "${demo_duration}"
  kill "${BLA_loading_animation_pid}" &> /dev/null
  clear
}

tput civis # Hide the terminal cursor
clear

# demo_loading_animation "${BLA_classic[@]}"
# demo_loading_animation "${BLA_box[@]}"
# demo_loading_animation "${BLA_bubble[@]}"
# demo_loading_animation "${BLA_breathe[@]}"
# demo_loading_animation "${BLA_growing_dots[@]}"
# demo_loading_animation "${BLA_passing_dots[@]}"
# demo_loading_animation "${BLA_metro[@]}"
demo_loading_animation "${BLA_snake[@]}"
# demo_loading_animation "${BLA_filling_bar[@]}"
# demo_loading_animation "${BLA_classic_utf8[@]}"
# demo_loading_animation "${BLA_bounce[@]}"
# demo_loading_animation "${BLA_vertical_block[@]}"
# demo_loading_animation "${BLA_horizontal_block[@]}"
# demo_loading_animation "${BLA_quarter[@]}"
# demo_loading_animation "${BLA_triangle[@]}"
# demo_loading_animation "${BLA_semi_circle[@]}"
# demo_loading_animation "${BLA_rotating_eyes[@]}"
# demo_loading_animation "${BLA_firework[@]}"
# demo_loading_animation "${BLA_braille[@]}"
# demo_loading_animation "${BLA_braille_whitespace[@]}"
# demo_loading_animation "${BLA_trigram[@]}"
# demo_loading_animation "${BLA_arrow[@]}"
# demo_loading_animation "${BLA_bouncing_ball[@]}"
# demo_loading_animation "${BLA_big_dot[@]}"
# demo_loading_animation "${BLA_modern_metro[@]}"
# demo_loading_animation "${BLA_pong[@]}"
# demo_loading_animation "${BLA_earth[@]}"
# demo_loading_animation "${BLA_clock[@]}"
# demo_loading_animation "${BLA_moon[@]}"
# demo_loading_animation "${BLA_orange_pulse[@]}"
# demo_loading_animation "${BLA_blue_pulse[@]}"
# demo_loading_animation "${BLA_football[@]}"
# demo_loading_animation "${BLA_blink[@]}"
# demo_loading_animation "${BLA_camera[@]}"
# demo_loading_animation "${BLA_sparkling_camera[@]}"
# demo_loading_animation "${BLA_sick[@]}"
# demo_loading_animation "${BLA_monkey[@]}"
# demo_loading_animation "${BLA_bomb[@]}"

tput cnorm # Restore the terminal cursor

exit 0
