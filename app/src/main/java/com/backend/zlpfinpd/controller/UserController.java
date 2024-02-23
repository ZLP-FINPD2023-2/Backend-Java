package com.backend.zlpfinpd.controller;

import com.backend.zlpfinpd.model.dto.UserDto;
import com.backend.zlpfinpd.model.dto.UserRegisterRequest;
import com.backend.zlpfinpd.service.UserService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.security.Principal;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1/user")
@Slf4j
public class UserController {
    private final UserService userService;

    @GetMapping()
    public ResponseEntity<UserDto> getUser(Principal principal) {
        UserDto user = userService.getUserByEmail(principal.getName());
        return ResponseEntity.ok(user);
    }

    @PutMapping()
    public ResponseEntity<UserDto> updateUser(@Valid @RequestBody UserRegisterRequest request, Principal principal) {
        UserDto user = userService.update(request, principal.getName());
        return ResponseEntity.ok(user);
    }
    @DeleteMapping()
    public ResponseEntity<?> deleteUser(Principal principal) {
        userService.delete(principal.getName());
        return ResponseEntity.ok().build();
    }
}
