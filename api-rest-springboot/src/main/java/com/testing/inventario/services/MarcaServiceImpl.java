package com.testing.inventario.services;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.dto.MarcaDTO;
import com.testing.inventario.entities.model.Categoria;
import com.testing.inventario.entities.model.Marca;
import com.testing.inventario.exceptions.ResourceNotFoundException;
import com.testing.inventario.repositories.CategoriaRepository;
import com.testing.inventario.repositories.MarcaRepository;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;


@Service
@RequiredArgsConstructor
@Slf4j
@Transactional
public class MarcaServiceImpl implements MarcaService {

    private final MarcaRepository marcaRepository;


    @Override
    public MarcaDTO getById(Long id) {
        Marca marca = marcaRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Marca con ID " + id + " no encontrada"));
        MarcaDTO marcaDTO = new MarcaDTO();
        marcaDTO.setId(marca.getId());
        marcaDTO.setNombre(marca.getNombre());
        return marcaDTO;
    }

    @Override
    public List<MarcaDTO> getList() {
        List<Marca> marcas = marcaRepository.findAll();
        List<MarcaDTO> marcaDTOList = new ArrayList<>();
        for (Marca marca : marcas) {
            MarcaDTO marcaDTO = new MarcaDTO();
            marcaDTO.setId(marca.getId());
            marcaDTO.setNombre(marca.getNombre());
            marcaDTOList.add(marcaDTO);
        }
        return marcaDTOList;
    }

    @Override
    public void create(MarcaDTO marcaDTO) {
        Marca marca = new Marca();
        marca.setNombre(marcaDTO.getNombre());
        marcaRepository.save(marca);
    }

    @Override
    public void update(Long id, MarcaDTO marcaDTO) {
        Optional<Marca> marcaOptional = marcaRepository.findById(id);
        if (marcaOptional.isPresent()) {
            Marca marca = marcaOptional.get();
            marca.setNombre(marcaDTO.getNombre());
            marcaRepository.save(marca);
        } else {
            throw new ResourceNotFoundException("Marca con ID " + id + " no encontrada");
        }
    }

    @Override
    public void delete(Long id) {
        Optional<Marca> marcaOptional = marcaRepository.findById(id);
        if (marcaOptional.isPresent()) {
            marcaRepository.deleteById(id);
        } else {
            throw new ResourceNotFoundException("Marca con ID " + id + " no encontrada");
        }
    }
}
