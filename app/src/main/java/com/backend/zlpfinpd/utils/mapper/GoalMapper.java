package com.backend.zlpfinpd.utils.mapper;

import com.backend.zlpfinpd.model.dto.GoalDto;
import com.backend.zlpfinpd.model.entity.Goal;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.ReportingPolicy;

@Mapper(componentModel = "spring", unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface GoalMapper {
    @Mapping(target = "userId", source = "user.id")
    GoalDto toDto(Goal goal);

    Goal toEntity(GoalDto goalDto);
}

