package tests.java;

import java.util.*;

import tests.java.dir1.BaseClass;
import tests.java.dir1.dir2.InterfaceA;
import tests.java.dir1.dir2.InterfaceB;

import java.io.*;

/**
 * Comprehensive test class with all possible Java combinations
 */
public class TestClass extends BaseClass implements InterfaceA, InterfaceB {

    // ============================================
    // FIELD DECLARATIONS - All Access Modifiers
    // ============================================

    // Public fields
    public String publicField;
    public static String publicStaticField;
    public final String publicFinalField = "constant";
    public static final String PUBLIC_STATIC_FINAL = "CONSTANT";
    public volatile int publicVolatileField;
    public transient String publicTransientField;

    // Private fields
    private String privateField;
    private static String privateStaticField;
    private final String privateFinalField = "constant";
    private static final String PRIVATE_STATIC_FINAL = "CONSTANT";
    private volatile int privateVolatileField;
    private transient String privateTransientField;

    // Protected fields
    protected String protectedField;
    protected static String protectedStaticField;
    protected final String protectedFinalField = "constant";
    protected static final String PROTECTED_STATIC_FINAL = "CONSTANT";

    // Package-private (default) fields
    String packagePrivateField;
    static String packagePrivateStaticField;
    final String packagePrivateFinalField = "constant";

    // Generic fields
    public List<String> genericList;
    private Map<String, Integer> genericMap;
    protected Set<List<String>> nestedGeneric;
    private Map<String, List<Integer>> complexGeneric;

    // Array fields
    public String[] stringArray;
    private int[][] twoDimensionalArray;
    protected List<String>[] genericArray;

    // Primitive fields
    public int intField;
    private long longField;
    protected double doubleField;
    boolean booleanField;
    byte byteField;
    short shortField;
    float floatField;
    char charField;

    // ============================================
    // CONSTRUCTORS
    // ============================================

    public TestClass() {
    }

    public TestClass(String arg) {
        this();
    }

    private TestClass(int arg1, String arg2) {
        this();
    }

    protected TestClass(List<String> list) {
        this();
    }

    TestClass(Map<String, Integer> map, boolean flag) {
        this();
    }

    // ============================================
    // PUBLIC METHODS
    // ============================================

    // Basic public methods
    public void publicVoidMethod() {
    }

    public String publicStringMethod() {
        return "";
    }

    public int publicIntMethod(String arg) {
        return 0;
    }

    public List<String> publicGenericMethod(Map<String, Integer> map) {
        return new ArrayList<>();
    }

    // Public static methods
    public static void publicStaticVoidMethod() {
    }

    public static String publicStaticStringMethod(int arg) {
        return "";
    }

    public static <T> List<T> publicStaticGenericMethod(T item) {
        return new ArrayList<>();
    }

    // Public final methods
    public final void publicFinalMethod() {
    }

    public final String publicFinalStringMethod(String arg) {
        return arg;
    }

    // Public synchronized methods
    public synchronized void publicSynchronizedMethod() {
    }

    public synchronized String publicSynchronizedStringMethod() {
        return "";
    }

    // Public abstract methods (if this were abstract class)
    // public abstract void publicAbstractMethod();

    // Public native method
    public native void publicNativeMethod();

    // Public strictfp method
    public strictfp double publicStrictfpMethod(double a, double b) {
        return a + b;
    }

    // ============================================
    // PRIVATE METHODS
    // ============================================

    private void privateVoidMethod() {
    }

    private String privateStringMethod() {
        return "";
    }

    private int privateIntMethod(String arg1, int arg2) {
        return 0;
    }

    private List<String> privateGenericMethod() {
        return new ArrayList<>();
    }

    private static void privateStaticMethod() {
    }

    private static String privateStaticStringMethod(String arg) {
        return arg;
    }

    private final void privateFinalMethod() {
    }

    private synchronized void privateSynchronizedMethod() {
    }

    private Map<String, List<Integer>> privateComplexGenericMethod(
            String arg1,
            List<Integer> arg2,
            Map<String, String> arg3) {
        return new HashMap<>();
    }

    // ============================================
    // PROTECTED METHODS
    // ============================================

    protected void protectedVoidMethod() {
    }

    protected String protectedStringMethod() {
        return "";
    }

    protected int protectedIntMethod(String arg) {
        return 0;
    }

    protected static void protectedStaticMethod() {
    }

    protected static String protectedStaticStringMethod() {
        return "";
    }

    protected final void protectedFinalMethod() {
    }

    protected synchronized void protectedSynchronizedMethod() {
    }

    protected <T> T protectedGenericTypeMethod(T item) {
        return item;
    }

    protected <K, V> Map<K, V> protectedMultiGenericMethod(K key, V value) {
        return new HashMap<>();
    }

    // ============================================
    // PACKAGE-PRIVATE (DEFAULT) METHODS
    // ============================================

    void packagePrivateVoidMethod() {
    }

    String packagePrivateStringMethod() {
        return "";
    }

    int packagePrivateIntMethod(String arg1, int arg2, boolean arg3) {
        return 0;
    }

    static void packagePrivateStaticMethod() {
    }

    static String packagePrivateStaticStringMethod() {
        return "";
    }

    final void packagePrivateFinalMethod() {
    }

    synchronized void packagePrivateSynchronizedMethod() {
    }

    // ============================================
    // METHODS WITH VARIOUS PARAMETER TYPES
    // ============================================

    public void methodWithPrimitives(
            int i,
            long l,
            double d,
            boolean b,
            char c,
            byte bt,
            short s,
            float f) {
    }

    public void methodWithArrays(
            String[] arr1,
            int[] arr2,
            List<String>[] arr3) {
    }

    public void methodWithVarargs(String... args) {
    }

    public void methodWithGenericVarargs(List<String>... lists) {
    }

    public <T> void methodWithGenericType(T item) {
    }

    public <T extends Comparable<T>> void methodWithBoundedGeneric(T item) {
    }

    public <T extends List<String> & Serializable> void methodWithMultipleBounds(T item) {
    }

    public void methodWithWildcard(List<?> list) {
    }

    public void methodWithUpperBoundWildcard(List<? extends Number> list) {
    }

    public void methodWithLowerBoundWildcard(List<? super Integer> list) {
    }

    // ============================================
    // METHODS WITH EXCEPTIONS
    // ============================================

    public void methodThrowsException() throws Exception {
    }

    public void methodThrowsMultipleExceptions()
            throws IOException, IllegalArgumentException, RuntimeException {
    }

    public String methodWithReturnAndThrows(String arg) throws Exception {
        return arg;
    }

    // ============================================
    // METHODS WITH ANNOTATIONS
    // ============================================

    @Override
    public String toString() {
        return "TestClass";
    }

    @Deprecated
    public void deprecatedMethod() {
    }

    @SuppressWarnings("unchecked")
    public void rawTypeMethod() {
    }

    @SafeVarargs
    public final <T> void safeVarargsMethod(List<T>... lists) {
    }

    // ============================================
    // COMPLEX COMBINED MODIFIERS
    // ============================================

    public static final synchronized void publicStaticFinalSynchronizedMethod() {
    }

    private static final void privateStaticFinalMethod() {
    }

    protected static synchronized void protectedStaticSynchronizedMethod() {
    }

    public final synchronized String publicFinalSynchronizedMethodWithReturn() {
        return "";
    }

    // ============================================
    // NESTED CLASSES
    // ============================================

    public static class PublicStaticNestedClass {
        public void nestedMethod() {
        }
    }

    private static class PrivateStaticNestedClass {
        private void nestedMethod() {
        }
    }

    protected class ProtectedInnerClass {
        protected void innerMethod() {
        }
    }

    class PackagePrivateInnerClass {
        void innerMethod() {
        }
    }

    // ============================================
    // ENUM
    // ============================================

    // public enum PublicEnum {
    // VALUE1, VALUE2, VALUE3;

    // public void enumMethod() {
    // }
    // }

    // private enum PrivateEnum {
    // ITEM1, ITEM2
    // }

    // ============================================
    // INTERFACE
    // ============================================

    public interface PublicInnerInterface {
        void interfaceMethod();

        default void defaultMethod() {
        }

        static void staticInterfaceMethod() {
        }
    }
}