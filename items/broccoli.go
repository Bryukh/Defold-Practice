components {
  id: "pickup"
  component: "/items/green_particles.particlefx"
}
components {
  id: "script"
  component: "/items/item.script"
  properties {
    id: "score"
    value: "-250.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "speed"
    value: "600.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "item_animation"
  component: "/items/item_animation.script"
  properties {
    id: "rotation_max"
    value: "20.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "tint_duration"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "rotation_duration"
    value: "0.5"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 32.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"broccoli\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 128.0\n"
  "  y: 128.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/food.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
}
