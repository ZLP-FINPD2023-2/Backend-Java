package com.backend.zlpfinpd.service;

import com.backend.zlpfinpd.model.dto.GoalDto;
import com.backend.zlpfinpd.model.entity.User;
import com.backend.zlpfinpd.repository.GoalRepository;
import com.backend.zlpfinpd.utils.mapper.GoalMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class GoalService {
    private final GoalRepository goalRepository;
    private final UserService userService;
    private final GoalMapper goalMapper;

    public List<GoalDto> getGoals(String email) {
        Long userId = userService.getUserIdByEmail(email);
        var goals = goalRepository.getGoalsByUserId(userId);
        return goals.stream().map(goalMapper::toDto).toList();
    }

    public GoalDto createGoal(GoalDto goalDto, String email) {
        Long userId = userService.getUserIdByEmail(email);
        var goal = goalMapper.toEntity(goalDto);
        goal.setUser(User.builder().id(userId).build());
        return goalMapper.toDto(goalRepository.save(goal));
    }
}