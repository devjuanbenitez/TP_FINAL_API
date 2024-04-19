package com.testing.inventario.services;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.dto.MarcaDTO;
import com.testing.inventario.entities.dto.ProductoDTO;

import java.util.List;

public interface ProductoService {

    ProductoDTO getById(Long id);
    List<ProductoDTO> getList();
    void create(ProductoDTO productoDTO);

    void update(Long id, ProductoDTO productoDTO);

    void delete(Long id);

}
