package com.testing.inventario.services;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.dto.MarcaDTO;

import java.util.List;

public interface MarcaService {

    MarcaDTO getById(Long id);

    List<MarcaDTO> getList();
    void create(MarcaDTO marcaRequest);

    void update(Long id, MarcaDTO marcaRequest);

    void delete(Long id);

}
