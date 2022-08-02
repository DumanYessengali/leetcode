public class Main {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(
                4,
                new TreeNode(
                        2,
                        new TreeNode(1),
                        new TreeNode()),
                new TreeNode(
                        2,
                        new TreeNode(1),
                        new TreeNode()
                )
        );
        SymmetricTree s = new SymmetricTree();
        System.out.println(s.isSymmetric(root));
    }
}
