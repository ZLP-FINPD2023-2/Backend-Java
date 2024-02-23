package com.backend.zlpfinpd.model.dto;

import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@EqualsAndHashCode(callSuper = true)
@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserRegisterRequest extends UserDto {
    @Size(min = 8, message = "Password must be at least 8 characters long")
    private String password;
}
