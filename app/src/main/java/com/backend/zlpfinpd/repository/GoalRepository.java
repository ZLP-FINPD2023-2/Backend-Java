package com.backend.zlpfinpd.repository;

import com.backend.zlpfinpd.model.entity.Goal;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface GoalRepository extends JpaRepository<Goal, Long> {
    List<Goal> getGoalsByUserId(Long user_id);
}
