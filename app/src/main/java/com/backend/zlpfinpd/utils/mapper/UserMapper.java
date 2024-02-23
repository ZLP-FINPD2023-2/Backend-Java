package com.backend.zlpfinpd.utils.mapper;

import com.backend.zlpfinpd.model.dto.UserDto;
import com.backend.zlpfinpd.model.dto.UserRegisterRequest;
import com.backend.zlpfinpd.model.entity.User;
import org.mapstruct.Mapper;
import org.mapstruct.ReportingPolicy;

@Mapper(componentModel = "spring", unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface UserMapper {
    UserDto toDto(User user);

    User toEntity(UserRegisterRequest userDto);
}
