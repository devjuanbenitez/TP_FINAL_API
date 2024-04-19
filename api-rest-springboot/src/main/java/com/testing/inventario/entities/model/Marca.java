package com.testing.inventario.entities.model;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name="marca")
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class Marca {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @Column(nullable = false)
    private String nombre;
}
