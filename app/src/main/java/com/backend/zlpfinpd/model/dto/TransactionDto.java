package com.backend.zlpfinpd.model.dto;

import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigDecimal;
import java.time.LocalDateTime;

@Data
@AllArgsConstructor
@Builder
@NoArgsConstructor
public class TransactionDto {
    private Long id;
    private Long userId;
    private String title;
    @Builder.Default
    private LocalDateTime date = LocalDateTime.now();
    private BigDecimal amount;
    @NotNull
    private Long budgetFromId;
    @NotNull
    private Long budgetToId;
}
