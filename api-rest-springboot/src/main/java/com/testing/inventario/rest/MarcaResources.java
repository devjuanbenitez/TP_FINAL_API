package com.testing.inventario.rest;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.dto.MarcaDTO;
import com.testing.inventario.services.CategoriaService;
import com.testing.inventario.services.MarcaService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value ="${base.api}/marcas",produces = MediaType.APPLICATION_JSON_VALUE)
@RequiredArgsConstructor
public class MarcaResources {

    private final MarcaService marcaService;

    @PostMapping
    public ResponseEntity<String> createMarca(
            @RequestBody final MarcaDTO categoria) {
        marcaService.create(categoria);
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @GetMapping("/{id}")
    public ResponseEntity<MarcaDTO> getMarcaById(
            @PathVariable(name = "id") final Long id) {
        return ResponseEntity.ok(marcaService.getById(id));
    }
    @GetMapping()
    public List<MarcaDTO> getAllMarca() {
        return marcaService.getList();
    }

    @PutMapping("/{id}")
    public ResponseEntity<Void> updateMarca(
            @PathVariable(name = "id") final Long id,
            @RequestBody MarcaDTO marca) {
        marcaService.update(id, marca);
        return ResponseEntity.ok().build();
    }
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteMarca(
            @PathVariable(name = "id") final Long id) {
        marcaService.delete(id);
        return ResponseEntity.noContent().build();
    }

}
