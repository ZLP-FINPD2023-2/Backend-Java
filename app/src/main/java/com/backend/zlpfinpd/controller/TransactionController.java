package com.backend.zlpfinpd.controller;

import com.backend.zlpfinpd.model.dto.TransactionDto;
import com.backend.zlpfinpd.service.TransactionService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.security.Principal;
import java.time.LocalDate;
import java.util.List;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1/transaction")
public class TransactionController {
    private final TransactionService transactionService;

    @GetMapping()
    public ResponseEntity<List<TransactionDto>> getTransactions(@RequestParam LocalDate date_from, @RequestParam LocalDate date_to, Principal principal) {
        var transactions = transactionService.getTransactions(date_from, date_to, principal.getName());
        return ResponseEntity.ok(transactions);
    }

    @PostMapping
    public ResponseEntity<TransactionDto> createTransaction(@Valid @RequestBody TransactionDto transactionDto, Principal principal) {
        TransactionDto createdTransaction = transactionService.createTransaction(transactionDto, principal.getName());
        return ResponseEntity.status(HttpStatus.CREATED).body(createdTransaction);
    }
}
