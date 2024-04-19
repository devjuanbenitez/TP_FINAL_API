package com.testing.inventario.rest;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.model.Categoria;
import com.testing.inventario.services.CategoriaService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value ="${base.api}/categorias",produces = MediaType.APPLICATION_JSON_VALUE)
@RequiredArgsConstructor
public class CategoriaResources {

    private final CategoriaService categoriaService;

    @PostMapping
    public ResponseEntity<String> createCategoria(
            @RequestBody final CategoriaDTO categoria) {
        categoriaService.create(categoria);
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @GetMapping("/{id}")
    public ResponseEntity<CategoriaDTO> getCategoriaById(
            @PathVariable(name = "id") final Long id) {
        return ResponseEntity.ok(categoriaService.getById(id));
    }


    @GetMapping()
    public List<CategoriaDTO> getAllCategorias() {
        return categoriaService.getList();
    }

    @PutMapping("/{id}")
    public ResponseEntity<Void> updateCategoria(
            @PathVariable(name = "id") final Long id,
            @RequestBody CategoriaDTO categoria) {
        categoriaService.update(id, categoria);
        return ResponseEntity.ok().build();
    }
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteCategoria(
            @PathVariable(name = "id") final Long id) {
        categoriaService.delete(id);
        return ResponseEntity.noContent().build();
    }

}
