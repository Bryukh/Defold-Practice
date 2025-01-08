components {
  id: "move_left_script"
  component: "/background/move_left.script"
  properties {
    id: "speed"
    value: "80.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"clouds_3\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 1280.0\n"
  "  y: 720.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/backgrounds.atlas\"\n"
  "}\n"
  ""
}
