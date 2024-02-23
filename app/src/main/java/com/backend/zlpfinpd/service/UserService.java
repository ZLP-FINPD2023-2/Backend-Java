package com.backend.zlpfinpd.service;

import com.backend.zlpfinpd.model.dto.AuthenticationResponse;
import com.backend.zlpfinpd.model.dto.UserDto;
import com.backend.zlpfinpd.model.dto.UserLoginRequest;
import com.backend.zlpfinpd.model.dto.UserRegisterRequest;
import com.backend.zlpfinpd.model.entity.User;
import com.backend.zlpfinpd.repository.UserRepository;
import com.backend.zlpfinpd.utils.mapper.UserMapper;
import jakarta.persistence.EntityNotFoundException;
import lombok.RequiredArgsConstructor;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.lang.reflect.Field;
import java.util.Objects;

@Service
@RequiredArgsConstructor
public class UserService {
    private final UserRepository userRepository;
    private final UserMapper userMapper;
    private final PasswordEncoder passwordEncoder;
    private final JwtService jwtService;
    private final AuthenticationManager authenticationManager;

    public void delete(String email) {
        User user = getUser(email);
        userRepository.delete(user);
    }

    public AuthenticationResponse register(UserRegisterRequest request) {
        User user = userMapper.toEntity(request);
        user.setPassword(passwordEncoder.encode(request.getPassword()));
        userRepository.save(user);
        var token = jwtService.generateToken(user);
        return new AuthenticationResponse(token);
    }

    public AuthenticationResponse authenticate(UserLoginRequest request) {
        authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        request.getEmail(),
                        request.getPassword()

                )
        );
        User user = getUser(request.getEmail());
        String token = jwtService.generateToken(user);
        return new AuthenticationResponse(token);
    }

    public UserDto update(UserRegisterRequest request, String email) {
        User user = getUser(email);
        if (!Objects.equals(user.getId(), request.getId())) {
            throw new IllegalArgumentException("You cannot update user with id " + request.getId());
        }
        updateUser(user, userMapper.toEntity(request));
        userRepository.save(user);
        return userMapper.toDto(user);
    }

    public UserDto getUserByEmail(String email) {
        User user = getUser(email);
        return userMapper.toDto(user);
    }

    public User getUser(String email) {
        return userRepository.findByEmail(email).orElseThrow(
                () -> new EntityNotFoundException("User not found")
        );
    }

    public Long getUserIdByEmail(String email) {
        return userRepository.findByEmail(email)
                .orElseThrow(() -> new EntityNotFoundException("User not found"))
                .getId();
    }

    private void updateUser(User prevState, User newState) {
        Field[] fields = User.class.getDeclaredFields();
        for (Field field : fields) {
            field.setAccessible(true);
            try {
                Object updatedValue = field.get(newState);
                if (updatedValue != null) {
                    field.set(prevState, updatedValue);
                }
            } catch (IllegalAccessException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
