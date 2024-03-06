package com.backend.zlpfinpd.controller;

import com.backend.zlpfinpd.model.dto.BudgetDto;
import com.backend.zlpfinpd.service.BudgetService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.security.Principal;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1/budget")
public class BudgetController {
    private final BudgetService budgetService;

    @GetMapping("/{budgetId}")
    public ResponseEntity<BudgetDto> getBudget(@PathVariable Long budgetId, Principal principal) {
        BudgetDto budgetDTO = budgetService.getBudgetById(budgetId, principal.getName());
        return ResponseEntity.ok(budgetDTO);
    }

    @PostMapping
    public ResponseEntity<BudgetDto> createBudget(@Valid @RequestBody BudgetDto budgetDto, Principal principal) {
        BudgetDto createdBudget = budgetService.createBudget(budgetDto, principal.getName());
        return ResponseEntity.status(HttpStatus.CREATED).body(createdBudget);
    }
}
