package com.testing.inventario.entities.model;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name="categoria")
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class Categoria {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @Column(nullable = false)
    private String nombre;
}
