package transpile

import "github.com/abibby/abc/parser"

func transpileTypeDefNode(s statements, n *parser.TypeDefNode) error {
	_, err := s.WriteString("typedef ")
	if err != nil {
		return err
	}
	err = transpileNode(s, n.Type)
	if err != nil {
		return err
	}

	_, err = s.WriteString(" " + n.Name.Value + ";")
	if err != nil {
		return err
	}
	return nil
}
func transpileStructDefNode(s statements, n *parser.StructDefNode) error {
	_, err := s.WriteString("struct " + n.Name + " {")
	if err != nil {
		return err
	}

	sTab := s.Tab()
	for _, prop := range n.Props {
		_, err = sTab.WriteString("\n")
		if err != nil {
			return err
		}
		err = transpileNode(sTab, prop.Type)
		if err != nil {
			return err
		}
		_, err = sTab.WriteString(" ")
		if err != nil {
			return err
		}
		err = transpileNode(sTab, prop.Name)
		if err != nil {
			return err
		}
		_, err = sTab.WriteString(";")
		if err != nil {
			return err
		}
	}

	_, err = s.WriteString("\n}")
	if err != nil {
		return err
	}

	return nil
}
func transpileStructInitNode(s statements, n *parser.StructInitNode) error {
	_, err := s.WriteString("(" + n.Type.Value + ")")
	if err != nil {
		return err
	}
	_, err = s.WriteString("{")
	if err != nil {
		return err
	}

	sTab := s.Tab()
	for i, prop := range n.Props {
		if i > 0 {
			_, err = sTab.WriteString(",")
			if err != nil {
				return err
			}
		}

		_, err = sTab.WriteString("\n    ." + prop.Key + " = ")
		if err != nil {
			return err
		}
		err = transpileNode(sTab, prop.Value)
		if err != nil {
			return err
		}
	}

	_, err = s.WriteString("\n}")
	if err != nil {
		return err
	}

	return nil
}
func transpileBasicTypeNode(s statements, n *parser.BasicTypeNode) error {
	_, err := s.WriteString(n.Value)
	return err
}

func transpilePointerTypeNode(s statements, n *parser.PointerTypeNode) error {
	err := transpileNode(s, n.Type)
	if err != nil {
		return err
	}
	_, err = s.WriteString("*")
	if err != nil {
		return err
	}
	return err
}

func transpileNewNode(s statements, n *parser.NewNode) error {
	_, err := s.WriteString("new_pointer(sizeof(" + n.Value.GetType() + "), &")
	if err != nil {
		return err
	}

	err = transpileNode(s, n.Value)
	if err != nil {
		return err
	}

	_, err = s.WriteString(")")
	if err != nil {
		return err
	}
	return nil
}
