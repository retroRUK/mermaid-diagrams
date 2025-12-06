package tests.java.dir1.dir2;

import java.util.List;
import java.util.Map;

public interface InterfaceB {
    // ============================================
    // PUBLIC METHODS
    // ============================================

    // Basic public methods
    public void publicVoidMethod();

    public String publicStringMethod();

    public int publicIntMethod(String arg);

    public List<String> publicGenericMethod(Map<String, Integer> map);
}
