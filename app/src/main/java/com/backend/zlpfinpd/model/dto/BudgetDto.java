package com.backend.zlpfinpd.model.dto;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class BudgetDto {
    private Long id;
    private Long userId;
    @NotNull
    private Long goalId;
    @NotBlank
    private String title;
}
