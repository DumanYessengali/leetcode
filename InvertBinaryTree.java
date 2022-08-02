public class InvertBinaryTree {
    public static TreeNode invertTree(TreeNode root) {
        inorder(root);
        return root;
    }

    public static void inorder(TreeNode root) {
        if (root == null) {
            return;
        }
        swap(root);
        inorder(root.left);
        inorder(root.right);
    }

    public static void swap(TreeNode root) {
        if (root == null) {
            return;
        }

        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode(
                4,
                new TreeNode(
                        2,
                        new TreeNode(1),
                        new TreeNode(3)),
                new TreeNode(
                        7,
                        new TreeNode(6),
                        new TreeNode(9)
                )
        );
        invertTree(root);
        TreeNode.print(root);
    }


}
