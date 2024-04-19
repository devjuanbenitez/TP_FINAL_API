package com.testing.inventario.services;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.dto.MarcaDTO;
import com.testing.inventario.entities.dto.ProductoDTO;
import com.testing.inventario.entities.model.Categoria;
import com.testing.inventario.entities.model.Marca;
import com.testing.inventario.entities.model.Producto;
import com.testing.inventario.exceptions.ResourceNotFoundException;
import com.testing.inventario.repositories.CategoriaRepository;
import com.testing.inventario.repositories.MarcaRepository;
import com.testing.inventario.repositories.ProductoRepository;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;


@Service
@RequiredArgsConstructor
@Slf4j
@Transactional
public class ProductoServiceImpl implements ProductoService {

    private final ProductoRepository productoRepository;
    @Autowired
    private MarcaRepository marcaRepository;
    @Autowired
    private CategoriaRepository categoriaRepository;


    @Override
    public ProductoDTO getById(Long id) {
        Producto producto = productoRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Producto con ID " + id + " no encontrada"));
        ProductoDTO productoDTO = new ProductoDTO();
        productoDTO.setId(producto.getId());
        productoDTO.setNombre(producto.getNombre());
        productoDTO.setPrecio(producto.getPrecio());
        productoDTO.setStock(producto.getStock());
        productoDTO.setCategoria(producto.getCategoria().getId());
        productoDTO.setMarca(producto.getMarca().getId());
        return productoDTO;
    }

    @Override
    public List<ProductoDTO> getList() {
        List<Producto> productos = productoRepository.findAll();
        List<ProductoDTO> productoDTOList = new ArrayList<>();
        for (Producto producto : productos) {
            ProductoDTO productoDTO = new ProductoDTO();
            productoDTO.setId(producto.getId());
            productoDTO.setNombre(producto.getNombre());
            productoDTO.setPrecio(producto.getPrecio());
            productoDTO.setStock(producto.getStock());
            if (producto.getCategoria() != null) {
                productoDTO.setCategoria(producto.getCategoria().getId());
            }
            if (producto.getMarca() != null) {
                productoDTO.setMarca(producto.getMarca().getId());
            }
            productoDTOList.add(productoDTO);
        }
        return productoDTOList;
    }



    public void create(ProductoDTO productoDTO) {
        // Verificar si la marca existe
        Marca marca = marcaRepository.findById(productoDTO.getMarca())
                .orElseThrow(() -> new ResourceNotFoundException("Marca con ID " + productoDTO.getMarca() + " no encontrada"));

        // Verificar si la categoría existe
        Categoria categoria = categoriaRepository.findById(productoDTO.getCategoria())
                .orElseThrow(() -> new ResourceNotFoundException("Categoría con ID " + productoDTO.getCategoria() + " no encontrada"));
        Producto producto = new Producto();
        producto.setNombre(productoDTO.getNombre());
        producto.setPrecio(productoDTO.getPrecio());
        producto.setStock(productoDTO.getStock());
        producto.setCategoria(categoria);
        producto.setMarca(marca);
        productoRepository.save(producto);
    }

    public void update(Long id, ProductoDTO productoDTO) {
        // Busca el producto en base al ID proporcionado
        Producto producto = productoRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Producto con ID " + id + " no encontrado"));
        // Actualiza los atributos del producto con los valores proporcionados
        producto.setNombre(productoDTO.getNombre());
        producto.setPrecio(productoDTO.getPrecio());
        producto.setStock(productoDTO.getStock());
        Marca marca = marcaRepository.findById(productoDTO.getMarca())
                .orElseThrow(() -> new ResourceNotFoundException("Marca con ID " + productoDTO.getMarca() + " no encontrada"));
        producto.setMarca(marca);
        Categoria categoria = categoriaRepository.findById(productoDTO.getCategoria())
                .orElseThrow(() -> new ResourceNotFoundException("Categoría con ID " + productoDTO.getCategoria() + " no encontrada"));
        producto.setCategoria(categoria);
        // Guarda el producto actualizado en el repositorio
        productoRepository.save(producto);
    }


    @Override
    public void delete(Long id) {
        Optional<Producto> productoOptional = productoRepository.findById(id);
        if (productoOptional.isPresent()) {
            productoRepository.deleteById(id);
        } else {
            throw new ResourceNotFoundException("Producto con ID " + id + " no encontrada");
        }
    }

}
