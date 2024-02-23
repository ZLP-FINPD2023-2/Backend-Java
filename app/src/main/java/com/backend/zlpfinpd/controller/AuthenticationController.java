package com.backend.zlpfinpd.controller;

import com.backend.zlpfinpd.model.dto.AuthenticationResponse;
import com.backend.zlpfinpd.model.dto.UserLoginRequest;
import com.backend.zlpfinpd.model.dto.UserRegisterRequest;
import com.backend.zlpfinpd.service.UserService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/auth")
@RequiredArgsConstructor
public class AuthenticationController {
    private final UserService userService;

    @PostMapping("/register")
    public ResponseEntity<AuthenticationResponse> register(@Valid @RequestBody UserRegisterRequest request) {
        return ResponseEntity.ok(userService.register(request));
    }

    @PostMapping("authenticate")
    public ResponseEntity<AuthenticationResponse> register(@Valid @RequestBody UserLoginRequest request) {
        return ResponseEntity.ok(userService.authenticate(request));
    }
}
