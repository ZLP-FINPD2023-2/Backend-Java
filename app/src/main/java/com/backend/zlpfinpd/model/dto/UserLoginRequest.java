package com.backend.zlpfinpd.model.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class UserLoginRequest {
    @NotNull
    @Email
    private String email;
    @NotNull
    @Size(min = 8, message = "Password must be at least 8 characters long")
    private String password;
}
