package com.backend.zlpfinpd.repository;

import com.backend.zlpfinpd.model.entity.Budget;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface BudgetRepository extends JpaRepository<Budget, Long> {
    Optional<Budget> getBudgetById(Long id);
}

