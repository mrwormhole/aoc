program day_2

    use str_list
    use rock_paper_scissors
    Use, intrinsic :: iso_fortran_env, Only : iostat_end
    implicit none

    character(len = *), parameter :: filename = "input.txt"
    character(len = 1), dimension(:), allocatable :: enemy_turns, player_turns
    character(len = 1):: enemy_turn, player_turn
    integer :: error, rounds, i, out_point, total_points

    rounds = 0
    open (10, file = filename)
    do
        read(10, *, iostat = error) enemy_turn, player_turn
        select case(error)
        case(0)
            call add_to_list(enemy_turns, enemy_turn)
            call add_to_list(player_turns, player_turn)
            rounds = rounds + 1
        case(iostat_end)
           exit
        case default
           print *, 'failed to read the file, error code:', error
           stop
        end select
    end do

    do i=1, rounds
        call calculate_points(enemy_turns(i), player_turns(i), out_point)
        total_points = total_points + out_point
    end do
    print *, "round 1:", total_points

    total_points = 0
    do i=1, rounds
        call calculate_points_via_strategy(enemy_turns(i), player_turns(i), out_point)
        total_points = total_points + out_point
    end do

    print *, "round 2:", total_points
    if (rounds > 0) then
        deallocate (enemy_turns)
        deallocate (player_turns)
    end if

end program day_2

module str_list
    contains

    subroutine add_to_list(list, element)
        
        implicit none

        integer :: i, isize
        character, intent(in) :: element
        character, dimension(:), allocatable, intent(inout) :: list
        character, dimension(:), allocatable :: temp_list

        if(allocated(list)) then
            isize = size(list)
            allocate(temp_list(isize+1))
            do i=1,isize          
                temp_list(i) = list(i)
            end do
            temp_list(isize+1) = element
            deallocate(list)
            call move_alloc(temp_list, list)
        else
            allocate(list(1))
            list(1) = element
        end if

    end subroutine add_to_list
end module str_list

module rock_paper_scissors

    implicit none
    
    character(len=1), dimension(3) :: enemy_choices = [character(len=1) :: "A", "B", "C"]         
    character(len=1), dimension(3) :: player_choices = [character(len=1) :: "X", "Y", "Z"]        

    contains

    subroutine calculate_points(enemy_choice, player_choice, points)
        character(len = 1), intent(in) :: enemy_choice, player_choice
        integer, intent (out) :: points
        integer :: player_index, enemy_index

        player_index = findloc(player_choices, player_choice, 1) - 1
        enemy_index = findloc(enemy_choices, enemy_choice, 1) - 1
        points =  player_index + 1
        if (player_index == enemy_index) then
            points = points + 3 ! draw round
        else if (enemy_index /= mod(player_index + 1, 3)) then
            points = points + 6 ! won the round
        end if
    end subroutine calculate_points

    subroutine calculate_points_via_strategy(enemy_choice, strategy, points)
        character(len = 1), intent(in) :: enemy_choice, strategy
        integer, intent (out) :: points
        integer :: expected_player_index, enemy_index, strategy_index

        strategy_index = findloc(player_choices, strategy, 1) - 1
        enemy_index = findloc(enemy_choices, enemy_choice, 1) - 1
        if (strategy_index == 1) then
            expected_player_index = enemy_index ! draw the round
        else if (strategy_index == 0) then
            if (enemy_index /= 0) then
                expected_player_index = enemy_index - 1 ! lose the round
            else
                expected_player_index = 2
            end if
        else 
            expected_player_index = mod(enemy_index + 1, 3) ! win the round       
        end if

        call calculate_points(enemy_choices(enemy_index+1), player_choices(expected_player_index+1), points)
    end subroutine
    
end module rock_paper_scissors