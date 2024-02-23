package com.backend.zlpfinpd.repository;

import com.backend.zlpfinpd.model.entity.Transaction;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;

@Repository
public interface TransactionRepository extends JpaRepository<Transaction, Long> {
    @Query("SELECT t FROM Transaction t WHERE t.date >= :dateFrom AND t.date <= :dateTo AND t.user.id = :userId")
    List<Transaction> findTransactionsBetweenDates(LocalDateTime dateFrom, LocalDateTime dateTo, Long userId);
}
