# Cài đặt thư viện (nếu chưa cài đặt, bạn cần chạy lệnh này trên terminal)
# pip install sentence-transformers

from sentence_transformers import SentenceTransformer, util

# Khởi tạo mô hình


def generate_score_with_many_projects(student_text, enterprise_texts):
    """
    hàm này trả về list các công việc theo thứ tự giảm dần của độ tương đồng cosine giữa văn bản của sinh viên và văn bản của doanh nghiệp 
    """
    # Tính embedding cho từng văn bản
    model = SentenceTransformer('./matching/matching-model')
    embedding_student = model.encode(student_text, convert_to_tensor=True)
    cosine_similarity_list = [] 
    for enterprise_text in enterprise_texts: 
        embedding_enterprise = model.encode(enterprise_text, convert_to_tensor=True) 
        cosine_sim = util.cos_sim(embedding_student, embedding_enterprise).item()
        cosine_similarity_list.append(cosine_sim * 100) # Chuyển đổi sang phần trăm

    # Sắp xếp danh sách theo độ tương đồng giảm dần

    return cosine_similarity_list
# Ví dụ: nhập văn bản của sinh viên và doanh nghiệp

def generate_score(student_text, enterprise_text):
    model = SentenceTransformer('all-MiniLM-L6-v2')
    embedding_student = model.encode(student_text, convert_to_tensor=True)
    embedding_enterprise = model.encode(enterprise_text, convert_to_tensor=True)

    return util.cos_sim(embedding_student, embedding_enterprise).item() * 100

if __name__ == "__main__":


    


    # nguyện vọng của ứng viên / sinh viên | Dạng text | 

    student_aspiration = "I aspire to develop groundbreaking AI technology solutions and participate in innovative projects."


    # Yêu cầu của danh nghiệp | Dạng lish các cặp [ id của project , yêu cầu của doanh nghiệp ]
    enterprise_requirements = [ [1, "We are seeking students who are passionate about technology and possess creative thinking to implement artificial intelligence projects."] , 
                            [2, "We are seeking students who are passionate about technology and possess creative thinking to implement website projects."] ] 


    # Kiểm tra khả năng hợp tác
    print(generate_score(student_aspiration, enterprise_requirements[0][1]))