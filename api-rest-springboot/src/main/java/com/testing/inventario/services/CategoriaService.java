package com.testing.inventario.services;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.model.Categoria;

import java.util.List;

public interface CategoriaService {

    CategoriaDTO getById(Long id);

    List<CategoriaDTO> getList();
    void create(CategoriaDTO categoriaRequest);

    void update(Long id, CategoriaDTO categoriaRequest);

    void delete(Long id);

}
