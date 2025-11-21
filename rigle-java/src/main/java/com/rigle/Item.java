package com.rigle;

import io.quarkus.hibernate.orm.panache.PanacheEntity;
import jakarta.persistence.Entity;
import jakarta.persistence.Table;

@Entity
@Table(name = "item")
public class Item extends PanacheEntity {
    public String name;
    public String description;
}
