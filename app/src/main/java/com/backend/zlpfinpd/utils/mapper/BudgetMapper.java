package com.backend.zlpfinpd.utils.mapper;

import com.backend.zlpfinpd.model.dto.BudgetDto;
import com.backend.zlpfinpd.model.entity.Budget;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.ReportingPolicy;


@Mapper(componentModel = "spring", unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface BudgetMapper {
    @Mapping(target = "userId", source = "user.id")
    @Mapping(target = "goalId", source = "goal.id")
    BudgetDto toDto(Budget budget);

    Budget toEntity(BudgetDto budgetDto);
}