components {
  id: "script"
  component: "/player/player.script"
  properties {
    id: "min_y"
    value: "120.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "max_y"
    value: "620.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "max_x"
    value: "1080.0"
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
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 23.0\n"
  "      y: -17.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 105.969284\n"
  "  data: 44.827908\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"fly\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/dragon.atlas\"\n"
  "}\n"
  ""
  scale {
    x: -1.0
  }
}
