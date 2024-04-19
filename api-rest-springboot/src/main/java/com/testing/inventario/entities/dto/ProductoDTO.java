package com.testing.inventario.entities.dto;

import com.testing.inventario.entities.model.Categoria;
import com.testing.inventario.entities.model.Marca;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class ProductoDTO {
    private Long id;
    private String nombre;
    private Long marca;
    private Long categoria;
    private Double precio;
    private Integer stock;
    // Getters y setters
}
